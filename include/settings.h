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
   NSString *configFile;
}
@property (readonly, assign, nonatomic) int port;
@property (readwrite, assign) BOOL shutdown;
@property (readonly, copy, nonatomic) NSString *configFile;

+ (Settings *)sharedInstance;
- (int)parseArguments:(int)argc:(const char *[])argv;
- (int)loadConfigFile;
@end
