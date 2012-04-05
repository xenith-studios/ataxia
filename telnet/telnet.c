#include <stdio.h>
#include <stdlib.h>
#include "telnet.h"
#include "_cgo_export.h"


static const telnet_telopt_t telopts[] = {
	{ TELNET_TELOPT_ECHO,			TELNET_WONT, TELNET_DONT },
	{ TELNET_TELOPT_EOR,			TELNET_WILL, TELNET_DONT },
	{ TELNET_TELOPT_TTYPE,			TELNET_WONT, TELNET_DO   },
	{ TELNET_TELOPT_NEW_ENVIRON,	TELNET_WONT, TELNET_DO   },
	{ TELNET_TELOPT_COMPRESS2,		TELNET_WILL, TELNET_DONT },
	{ TELNET_TELOPT_ZMP,			TELNET_WILL, TELNET_DONT },
	{ -1, 0, 0 }
};

telnet_t *initialize()
{
	telnet_t *telnet = telnet_init(telopts, event_handler, 0, NULL);
	return telnet;
}

int cleanup(telnet_t *telnet)
{
	telnet_free(telnet);
	telnet = NULL;
	return 0;
}

void event_handler(telnet_t *telnet, telnet_event_t *event, void *user_data)
{
	switch (event->type) {
		case TELNET_EV_DATA:
			// Call the input-handling callback
			//process_user_input(user, event->data.buffer, event->data.size);
			break;
		case TELNET_EV_SEND:
			// Call the sending callback
			//write_to_descriptor(user, event->data.buffer, event->data.size);
			break;
		case TELNET_EV_ERROR:
			exit(1);
			//fatal_error("TELNET error: %s", event->error.msg);
			break;
	}
	return;
}
