//
//  settings.h
//  ataxia
//
//  Created by Justin Rocha on 12/31/09.
//  Copyright 2009 Xenith Studios. All rights reserved.
//

#import <Foundation/Foundation.h>


@interface Settings : NSObject {
   int port;
   BOOL shutdown;
}
@property int port;
@property BOOL shutdown;

+ (Settings *)sharedInstance;
- (void)parseArguments:(int)argc:(const char *[])argv;
- (void)loadConfigFile;
@end
