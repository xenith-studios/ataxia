/*
 * Ataxia Engine Project
 * Copyright (C) 2009  Xenith Studios
 * See COPYING for license details
 * http://github.com/xenith/ataxia/
 *
 */

#import <Foundation/Foundation.h>
#import "settings.h"

void sig_term_handler(int signum) { 
   if ((signum != SIGSEGV) && (signum != SIGBUS) && (signum != SIGIOT)) {
      NSLog(@"Got signal: %i,  NOT setting to default handler.", signum);
      signal(signum, (&sig_term_handler));
   } else {
      NSLog(@"Got signal: %i, setting to default handler.", signum);
      signal(signum, SIG_DFL);
   }

   switch (signum) {
      // Signals that we will shutdown on
      case SIGTERM:
         NSLog(@"Received SIGTERM, shutting down.");
         exit(0);
         break;

      case SIGINT:
         NSLog(@"Received SIGINT, shutting down.");
         exit(0);
         break;

      case SIGSTOP:
         NSLog(@"Receieved SIGSTOP, shutting down.");
         exit(0);
         break;

      case SIGXCPU:
         NSLog(@"Receieved SIGXCPU, shutting down.");
         exit(0);
         break;

      case SIGXFSZ:
         NSLog(@"Receieved SIGXFSZ, shutting down.");
         exit(0);
         break;		

      // Signals that we have crashed
      case SIGSEGV:
      case SIGIOT:
         NSLog(@"Received SIGSEV or SIGIOT, shutting down.");
         NSLog(@"Ataxia has crashed: %i", getpid());
         abort();
         break;

      case SIGBUS:
         NSLog(@"Received SIGBUS, shutting down.");
         NSLog(@"Ataxia has crashed: %i", getpid());
         abort();
         break;

      // Signals that we will ignore
      case SIGALRM:
         NSLog(@"Received SIGALRM, ignoring...");
         break;

      case SIGHUP:
         NSLog(@"Receieved SIGHUP, ignoring...");
         break;

      case SIGIO:
         NSLog(@"Receieved SIGIO, ignoring...");
         break;

      case SIGPIPE:
         NSLog(@"Receieved SIGPIPE, ignoring...");
         break;

      case SIGTTIN:
         NSLog(@"Receieved SIGTTIN, ignoring...");
         break;

      case SIGTTOU:
         NSLog(@"Receieved SIGTTOU, ignoring...");
         break;

      case SIGUSR1:
      case SIGUSR2:
         NSLog(@"Receieved SIGUSR*, ignoring...");
         break;

      case SIGVTALRM:
         NSLog(@"Receieved SIGVTARLM, ignoring...");
         break;

		default:
			NSLog(@"Receieved unknown signal, shutting down.");
			exit(0);
			break;
	}
}

int main (int argc, const char * argv[]) {
   NSAutoreleasePool * pool = [[NSAutoreleasePool alloc] init];
    
   // Settings - parse command-line and load config file
   [[Settings sharedInstance] loadConfigFile];
   [[Settings sharedInstance] parseArguments:argc :argv];

   // Message queue - initialize

   // Initialize and start comms server
    
   // Game loop - while !shutdown
      // Handle messages (select?)
      // Check for new connections
      // Game tick
      // Sleep

   [pool drain];
   return 0;
}
