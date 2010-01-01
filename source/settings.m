//
//  settings.m
//  ataxia
//
//  Created by Justin Rocha on 12/31/09.
//  Copyright 2009 Xenith Studios. All rights reserved.
//

#import "settings.h"

@implementation Settings

static Settings * sharedInstance = nil;

@synthesize port, shutdown;

#pragma mark -
#pragma mark class instance methods

- (id)init
{
   if (self = [super init]) {
      port = 0;
      shutdown = NO;
   }
   return self;
}

- (void)parseArguments:(int)argc:(const char *[])argv
{
   
}
- (void)loadConfigFile
{
   
}

#pragma mark -
#pragma mark Singleton methods

+ (void)initialize
{
   static BOOL initialized = NO;
   if(!initialized) {
      initialized = YES;
      sharedInstance = [[Settings alloc] init];
   }
}

+ (Settings *)instance
{
   if (sharedInstance)
      return sharedInstance;

   @synchronized(self) {
      if (sharedInstance == nil)
         sharedInstance = [[Settings alloc] init];
   }   
   return sharedInstance;
}

+ (id)allocWithZone:(NSZone *)zone
{
   @synchronized(self) {
      if (sharedInstance == nil) {
         sharedInstance = [super allocWithZone:zone];
         return sharedInstance;  // assignment and return on first allocation
      }
   }
   return nil; // on subsequent allocation attempts return nil
}

- (id)copyWithZone:(NSZone *)zone
{
   return self;
}

- (id)retain
{
   return self;
}

- (unsigned)retainCount
{
   return UINT_MAX;  // denotes an object that cannot be released
}

- (void)release
{
   //do nothing
}

- (id)autorelease
{
   return self;
}

@end