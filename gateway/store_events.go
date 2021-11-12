package gateway

func (g *Gateway) handleStoreEvent(iface interface{}) {
	switch ev := iface.(type) {
	case *ReadyEvent:
		StoreDebug("Received ready event with %v channels, %v servers, %v users", len(ev.Channels), len(ev.Servers), len(ev.Users))

		for _, ch := range ev.Channels {
			g.Store.PutChannel(&ch)
		}

		for _, s := range ev.Servers {
			g.Store.PutServer(&s)
		}

		for _, u := range ev.Users {
			g.Store.PutUser(&u)
		}

	case *ChannelCreateEvent:
		g.Store.PutChannel(&ev.Channel)

	case *ChannelUpdateEvent:
		ch, err := g.Store.Channel(ev.ID)
		if err != nil {
			// not much we can do here
			StoreDebug("Received channel update for %v but not in cache", ev.ID)
			return
		}

		ev.Update(ch)

		g.Store.PutChannel(ch)
	}
}
