//
//  log.m
//  ataxia
//
//  Created by Justin Rocha on 12/31/09.
//  Copyright 2009-2010 Xenith Studios. All rights reserved.
//

#import "log.h"

@implementation Log

static Log *sharedInstance = nil;

#pragma mark -
#pragma mark class instance methods

- (id)init
{
   if (self = [super init]) {
   }
   return self;
}

- (void)dealloc
{
   [super dealloc];
}

#pragma mark -
#pragma mark Singleton methods

+ (void)initialize
{
   static BOOL initialized = NO;
   if(!initialized) {
      initialized = YES;
      sharedInstance = [[Log alloc] init];
   }   
}

+ (Log *)sharedInstance
{
   if (sharedInstance)
      return sharedInstance;

   @synchronized(self) {
      if (sharedInstance == nil)
         sharedInstance = [[Log alloc] init];
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