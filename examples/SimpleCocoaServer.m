//  
//  SimpleCocoaServer, a basic server class written in objectiv-c for use in cocoa applications
//   -- v1.0 --
//   SimpleCocoaServer.m
//   ------------------------------------------------------
//  | Created by David J. Koster, release 26.08.2009.      |
//  | Copyright 2008 David J. Koster. All rights reserved. |
//  | http://www.david-koster.de/code/simpleserver         |
//  | code@david-koster.de for help or see:                |
//  | http://sourceforge.net/projects/simpleserver         |
//   ------------------------------------------------------
// 
//  This program is free software: you can redistribute it and/or modify
//  it under the terms of the GNU General Public License as published by
//  the Free Software Foundation, either version 3 of the License, or
//  (at your option) any later version.
// 
//  This program is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//  GNU General Public License for more details.
// 
//  You should have received a copy of the GNU General Public License
//  along with this program.  If not, see <http://www.gnu.org/licenses/>. */
//


#import "SimpleCocoaServer.h"
#import <sys/socket.h>
#import <netinet/in.h>
#import <arpa/inet.h>


@implementation SimpleCocoaServer

#pragma mark Class Methods

+ (id)server
{
	self = [[self alloc] init];
	return self;
}

+ (id)serverWithPort:(int)initPort delegate:(id)initDelegate
{
	self = [[self alloc] initWithPort:initPort delegate:initDelegate];
	return self;
}

#pragma mark Instance Methods

- (id)init
{
	if(self = [super init]) {
		
		//NSAssert(delegate != nil, @"Please specify a delegate");
		//NSAssert([delegate respondsToSelector:@selector(processMessage:connection:)],
		//		 @"Delegate needs to implement 'processMessage:connection:'");
		
		serverPort = 0;
		serverDelegate = nil;
		connections = [[NSMutableArray alloc] init];
		lAddr = SCSListenAll;
		//initialize lStrAddr;
		isListening = NO;
		
	}
	
	return self;	
}

- (id)initWithPort:(int)initPort delegate:(id)initDelegate
{
	if(self = [super init]) {
		
		//NSAssert(delegate != nil, @"Please specify a delegate");
		//NSAssert([delegate respondsToSelector:@selector(processMessage:connection:)],
		//		 @"Delegate needs to implement 'processMessage:connection:'");
		
		serverPort = initPort;
		serverDelegate = [initDelegate retain];
		connections = [[NSMutableArray alloc] init];
		lAddr = SCSListenAll;
		isListening = NO;
		
	}
	
	return self;
	
}

- (void)dealloc
{
	if(isListening) {
		[[NSNotificationCenter defaultCenter] removeObserver:self];
		[fileHandle release];
	}
	[connections release];
	[serverDelegate release];
	[super dealloc];
}

#pragma mark Listening Methods

- (SCSInit)startListening
{
	if(isListening)
		return SCSInitError_Listening;
	if(serverPort < 1)
		return SCSInitError_Port;
	if(!serverDelegate)
		return SCSInitError_Delegate;

	int filedescriptor = -1;
	CFSocketRef socket = CFSocketCreate(kCFAllocatorDefault, PF_INET, SOCK_STREAM, IPPROTO_TCP, 1, NULL, NULL);
	
	if(socket) {
		
		filedescriptor = CFSocketGetNative(socket);
		
		//this code prevents the socket from existing after the server has crashed or been forced to close
		
		int yes = 1;
		setsockopt(filedescriptor, SOL_SOCKET, SO_REUSEADDR, (void *)&yes, sizeof(yes));
		
		struct sockaddr_in addr4;
		memset(&addr4, 0, sizeof(addr4));
		addr4.sin_len = sizeof(addr4);
		addr4.sin_family = AF_INET;
		addr4.sin_port = htons(serverPort);
		if(lAddr == SCSListenLoopback)
			addr4.sin_addr.s_addr = htonl(INADDR_LOOPBACK);
		else if(lAddr == SCSListenLocal)
			addr4.sin_addr.s_addr = htonl(INADDR_LOOPBACK);
		else if((lAddr == SCSListenOther) && (lStrAddr))
			inet_pton(AF_INET, lStrAddr, &addr4.sin_addr);
		else
			addr4.sin_addr.s_addr = htonl(INADDR_ANY); //any network address, e.g. 127.0.0.1, 168.192.2.101 etc;
		// inet_pton(AF_INET, "127.0.0.1", &addr4.sin_addr); would set the listening IP to only 127.0.0.1
		NSData *address4 = [NSData dataWithBytes:&addr4 length:sizeof(addr4)];
		
		if (kCFSocketSuccess != CFSocketSetAddress(socket, (CFDataRef)address4))
			return SCSInitError_Bind;
		
	} else {
		return SCSInitError_NoSocket;
	}
	
	fileHandle = [[NSFileHandle alloc] initWithFileDescriptor:filedescriptor
											   closeOnDealloc:YES];
	NSNotificationCenter *nc = [NSNotificationCenter defaultCenter];
	[nc addObserver:self
		   selector:@selector(newConnection:)
			   name:NSFileHandleConnectionAcceptedNotification
			 object:nil];
	[fileHandle acceptConnectionInBackgroundAndNotify];
	
	isListening = YES;
	return SCSInitOK;
	
}

- (void)stopListening
{
	if(!isListening)
		return; //Server is not listening
	
	//close every connection:
	while ([connections count] != 0) {
		[self closeConnection:[connections objectAtIndex:0]];
	}
	
	[[NSNotificationCenter defaultCenter] removeObserver:self]; //don't handle new requests;
	CFSocketRef socket = CFSocketCreateWithNative(kCFAllocatorDefault,[fileHandle fileDescriptor],1,NULL,NULL);
	CFSocketInvalidate(socket);
	CFRelease(socket);
	[fileHandle release];
	
	//server is not running anymore
	isListening = NO;
}

#pragma mark Accessor Methods

- (BOOL)isListening
{
	return isListening;
}

- (BOOL)setServerPort:(int)newPort
{
	if(!isListening)
		serverPort = newPort;
	else
		return NO;
	return YES;
}

- (int)serverPort
{
	return serverPort;
}

- (BOOL)setServerDelegate:(id)newDelegate
{
	if(!isListening) {
		[serverDelegate release];
		serverDelegate = newDelegate;
	} else {
		return NO;
	}
	return YES;
}

- (SCSListenAddress)listenAddress
{
	return lAddr;
}

- (NSString *)listenAddressAsString
{
	return [NSString stringWithUTF8String:lStrAddr];
}

- (void)setListenAddress:(SCSListenAddress)newLAddr
{
	lAddr = newLAddr;
	NSString *tmpListenAddr;
	if(lAddr == SCSListenLoopback)
		tmpListenAddr = @"127.0.0.1";
	else if(lAddr == SCSListenLocal) 
		tmpListenAddr = @"127.0.0.1"; //currently
	else
		tmpListenAddr = @"0.0.0.0";
	strncpy(lStrAddr,[tmpListenAddr UTF8String],15);
}

- (BOOL)setListenAddressByString:(NSString *)newStrAddr
{
	if(!isListening) {
		[self setListenAddress:SCSListenOther]; //needs to be called
		strncpy(lStrAddr, [newStrAddr UTF8String], 15); //before address is copied here
	} else {
		return NO;
	}
	return YES;
}

#pragma mark Delegate Methods

- (void)processMessage:(NSString *)message orData:(NSData *)data fromConnection:(SimpleCocoaConnection *)con;
{
	//for compatibility to older versions that used this class
	if([serverDelegate respondsToSelector:@selector(processMessage:fromConnection:)] && ![serverDelegate respondsToSelector:@selector(processMessage:orData:fromConnection:)])
		[serverDelegate processMessage:message fromConnection:con];//[serverDelegate performSelector:@selector(processMessage:fromConnection:) withObject:message withObject:con];
	else if([serverDelegate respondsToSelector:@selector(processMessage:orData:fromConnection:)])
		[serverDelegate processMessage:message orData:data fromConnection:con];
}

- (void)processNewConnection:(SimpleCocoaConnection *)con
{
	if([serverDelegate respondsToSelector:@selector(processNewConnection:)])
		[serverDelegate processNewConnection:con];
}

- (void)processClosingConnection:(SimpleCocoaConnection *)con
{
	if([serverDelegate respondsToSelector:@selector(processClosingConnection:)])
		[serverDelegate processClosingConnection:con];
}

//Not used in this version. Perhaps used and documented in future versions.
/*#pragma mark Delegate Error Methods

- (void)processNewConnectionFileHandleError:(NSNumber *)error
{
	if([serverDelegate respondsToSelector:@selector(processNewConnectionFileHandleError:)])
		[serverDelegate processNewConnectionFileHandleError:error];
}

- (void)processDataSendingError:(NSException *)exception forConnection:(SimpleCocoaConnection *)con
{
	if([serverDelegate respondsToSelector:@selector(processDataSendingError:forConnection:)])
		[serverDelegate processDataSendingError:exception forConnection:con];
}
*/
#pragma mark Connections

- (NSArray *)connections
{
	return connections;
}

- (void)newConnection:(NSNotification *)notification
{
	NSDictionary *userInfo = [notification userInfo];
	NSFileHandle *remoteFileHandle = [userInfo objectForKey:
									  NSFileHandleNotificationFileHandleItem];
	NSNumber *errorNo = [userInfo objectForKey:@"NSFileHandleError"];
	if(errorNo) {
		//[self processNewConnectionFileHandleError:errorNo]; //Not used in this version. Perhaps used and documented in future versions.
		return;
	}
	
	[fileHandle acceptConnectionInBackgroundAndNotify];
	
	if(remoteFileHandle) {
		SimpleCocoaConnection *connection = [[SimpleCocoaConnection alloc] initWithFileHandle:remoteFileHandle delegate:self];
		if(connection) {
			NSIndexSet *insertedIndexes = [NSIndexSet indexSetWithIndex:[connections count]];
            [self willChange:NSKeyValueChangeInsertion
             valuesAtIndexes:insertedIndexes forKey:@"connections"];
            [connections addObject:connection];
            [self didChange:NSKeyValueChangeInsertion
			valuesAtIndexes:insertedIndexes forKey:@"connections"];
            [connection release];
			[self processNewConnection:connection];
		}
	}
}

- (void)closeConnection:(SimpleCocoaConnection *)con
{
	[self processClosingConnection:con];
	int connectionIndex = [connections indexOfObjectIdenticalTo:con];
    if(connectionIndex == NSNotFound)
		return;
	NSIndexSet *connectionIndexSet = [NSIndexSet indexSetWithIndex:connectionIndex];
    [self willChange:NSKeyValueChangeRemoval valuesAtIndexes:connectionIndexSet
              forKey:@"connections"];
    [connections removeObjectsAtIndexes:connectionIndexSet];
	//the connection was released when added to the array
    [self didChange:NSKeyValueChangeRemoval valuesAtIndexes:connectionIndexSet
             forKey:@"connections"];
}

#pragma mark Sending Data
- (BOOL)sendData:(NSData *)data toConnection:(SimpleCocoaConnection *)con
{
	NSFileHandle *remoteFileHandle = [con fileHandle];
	@try {
		//[NSThread detachNewThreadSelector:@selector(finalSendData:)	toTarget:self withObject:self];
        [remoteFileHandle writeData:data];
    }
    @catch (NSException *exception) {
		//[self processDataSendingError:exception forConnection:con]; //Not used in this version. Perhaps used and documented in future versions.
		return NO;
    }
	return YES;
}

- (BOOL)sendString:(NSString *)string toConnection:(SimpleCocoaConnection *)con
{
	return [self sendData:[string dataUsingEncoding:NSASCIIStringEncoding] toConnection:con];
}

- (void)sendDataToAll:(NSData *)data
{
	NSEnumerator *en = [connections objectEnumerator];
	SimpleCocoaConnection *con;
	
	while(con = [en nextObject])
		[self sendData:data toConnection:con];
}

- (void)sendStringToAll:(NSString *)string
{
	[self sendDataToAll:[string dataUsingEncoding:NSASCIIStringEncoding]];
}

@end



@interface SimpleCocoaConnection (PrivateMethods)

- (void)setRemoteAddress:(NSString *)newAddress;
- (void)setRemotePort:(int)newPort;

@end

@implementation SimpleCocoaConnection

- (id)initWithFileHandle:(NSFileHandle *)fh delegate:(id)initDelegate
{
    if(self = [super init]) {
		fileHandle = [fh retain];
		connectionDelegate = [initDelegate retain];
		
		// Get IP address of remote client
		CFSocketRef socket = CFSocketCreateWithNative(kCFAllocatorDefault, [fileHandle fileDescriptor], kCFSocketNoCallBack, NULL, NULL);
		CFDataRef addrData = CFSocketCopyPeerAddress(socket);
		CFRelease(socket);
		
		if(addrData) {
			struct sockaddr_in *sock = (struct sockaddr_in *)CFDataGetBytePtr(addrData);
			[self setRemotePort:(sock->sin_port)];
			char *naddr = inet_ntoa(sock->sin_addr);
			[self setRemoteAddress:[NSString stringWithCString:naddr]];
			CFRelease(addrData);
		} else {
			[self setRemoteAddress:@"NULL"];
		}
		
		// Register for notification when data arrives
		NSNotificationCenter *nc = [NSNotificationCenter defaultCenter];
		[nc addObserver:self
			   selector:@selector(dataReceivedNotification:)
				   name:NSFileHandleReadCompletionNotification
				 object:fileHandle];
		[fileHandle readInBackgroundAndNotify];
	}
	return self;
}

- (void)dealloc
{
	[[NSNotificationCenter defaultCenter] removeObserver:self];
	[connectionDelegate release];
	[fileHandle closeFile];
	[fileHandle release];
	[remoteAddress release];
	[super dealloc];
}

- (NSString *)description
{
	return [NSString stringWithFormat:@"%@:%d",[self remoteAddress],[self remotePort]];
}

#pragma mark Accessor Methods

- (NSFileHandle *)fileHandle 
{
	return fileHandle;
}

- (void)setRemoteAddress:(NSString *)newAddress
{
	[remoteAddress release];
	remoteAddress = [newAddress copy];
}

- (NSString *)remoteAddress
{
	return remoteAddress;
}

- (void)setRemotePort:(int)newPort
{
	remotePort = newPort;
}

- (int)remotePort
{
	return remotePort;
}

#pragma mark Notification Methods

- (void)dataReceivedNotification:(NSNotification *)notification
{
	NSData *data = [[notification userInfo] objectForKey:NSFileHandleNotificationDataItem];
	
	if ([data length] == 0) {
		// NSFileHandle's way of telling us that the client closed the connection
		[connectionDelegate closeConnection:self];
		return;
	} else {
		[fileHandle readInBackgroundAndNotify];
		NSString *received = [[NSString alloc] initWithData:data encoding:NSASCIIStringEncoding];
		if([received characterAtIndex:0] == 0x04) { // End-Of-Transmission sent by client
			return;
		}
		[connectionDelegate processMessage:received orData:data fromConnection:self];
	}
}

@end