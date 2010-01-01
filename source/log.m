//
//  log.m
//  ataxia
//
//  Created by Justin Rocha on 12/31/09.
//  Copyright 2009 Xenith Studios. All rights reserved.
//

#import "log.h"


@implementation Log

static Log * instance = nil;

#pragma mark -
#pragma mark class instance methods

#pragma mark -
#pragma mark Singleton methods

+ (void)initialize
{
   
}

+ (Log *)instance
{
   @synchronized(self) {
      if (instance == nil)
         instance = [[Log alloc] init];
   }
   
   // Default values
   
   return instance;
}

+ (id)allocWithZone:(NSZone *)zone
{
   if (instance)
      return instance;

   @synchronized(self) {
      if (instance == nil) {
         instance = [super allocWithZone:zone];
         return instance;  // assignment and return on first allocation
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