//
//  network.h
//  ataxia
//
//  Created by Justin Rocha on 12/31/09.
//  Copyright 2009-2010 Xenith Studios. All rights reserved.
//

#import "AsyncSocket.h"

@interface Network : NSObject {
   AsyncSocket *listenSocket;
   NSMutableArray *connectedSockets;
}

@end