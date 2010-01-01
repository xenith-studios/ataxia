//
//  log.h
//  ataxia
//
//  Created by Justin Rocha on 12/31/09.
//  Copyright 2009 Xenith Studios. All rights reserved.
//

#import <Foundation/Foundation.h>

typedef enum {
    LOG_INFO,
    LOG_WARNING,
    LOG_ERROR,
    LOG_NETWORK,
    LOG_ZMP,
    LOG_ADMIN,
} LogLevel;

@interface Log : NSObject {
   LogLevel level;
}

+ (Log *)instance;
+ (void)initialize;
@end
