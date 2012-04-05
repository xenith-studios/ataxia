#include "libtelnet.h"

telnet_t *initialize();

int cleanup(telnet_t *telnet);

void event_handler(telnet_t *telnet, telnet_event_t *event, void *user_data);