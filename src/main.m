/*
 * Ataxia Engine Project
 * Copyright (C) 2009  Xenith Studios
 * See COPYING for license details
 * http://github.com/xenith/ataxia/
 *
 */

#import <Foundation/Foundation.h>
#import "settings.h"

void signal_handler(int signum) { 
   if ((signum != SIGSEGV) && (signum != SIGBUS) && (signum != SIGIOT)) {
      NSLog(@"Got signal: %i,  NOT setting to default handler.", signum);
      signal(signum, (&signal_handler));
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

   // Print program information
   printf("Ataxia Engine V0.1 © 2009-2010, Xenith Studios (see AUTHORS)\n"
            "Ataxia Engine comes with ABSOLUTELY NO WARRANTY; see COPYING for details.\n"
            "This is free software, and you are welcome to redistribute it\n"
            "under certain conditions; for details, see the file COPYING.\n\n");
   
   // Attempt to parse command line and config file. Exits on error.
   if ([[Settings sharedInstance] parseArguments:argc :argv]) {
      fprintf(stderr, "Usage: ataxia [CONFIGFILE] [OPTION]...\n\n"
      "\n");
      return 1;
   }
   if ([[Settings sharedInstance] loadConfigFile]) {
      NSLog(@"Failed to load config file: ", [[Settings sharedInstance] configFile]);
      return 1;
   }

   // Initializations
      // Logging
      // Driver
         // comms
         // chroot
      // Queues (messages, events, etc)
      // Lua
    
   // Game loop - while !shutdown
      // Handle network messages (push user events) -- NOT NEEDED, HANDLED AUTOMATICALLY BY NETWORKING/SOCKET CLASSES
      // Handle game updates
         // Game tick
         // Time update
         // Weather update
         // Entity updates (push events)
      // Handle pending events
      // Handle pending messages (network and player)
      // Sleep

   // Cleanup
   
   [pool drain];
   return 0;
}
