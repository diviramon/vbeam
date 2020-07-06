# Voronoibeam

The decentralized content delivery based on IPFS for internet service providers (ISPs).

## Overview

- *Vbeam pinpub*: a publisher announces to the world the list of data object ids (the pinlist) and regularly publishes changes or updates to it; the ISPs that are interested in keeping track of the pinlists for a given publisher are subscribing to such announcements; we can see that generally this sounds as a pubsub model.

- *Vbeam gateway*: a server to be run within the ISP infrastructure  that is largely a regular IPFS node. The ISP can subscribe their gateway to the publisher's pinlists and which adds the data object ids listed there to its own pinlist; the server also serves the acquired data objects over the HTTP gateway so that the ISP's customers can be served these data objects over HTTP.

- *Vbeam routing table*: the voronoibeam gateways go up and down all the time, subscribe and unsubscribe to pinlists and so on.  The ISPs are the ones who know best which of the gateways the traffic from their customers needs to redirected to, and they publish these rerouting recommendations within the decentralized routing table system.

- *Vbeam HTTP redirection*: the publisher can keep track of the routing tables for their peerlists for all the ISPs and redirect the customers's requests to the gateway recommended by their ISP.

## Specifications

- *Vbeam pinpub*: we can start with something RSS-style in spirit; that is, serving the pinlists as JSONs over HTTP and stuff; maybe a magic word DNS record that publishers can add to signify a given (workload estimate: 20 hours).

- *Vbeam gateway*:  thanks to libp2p by the Protocol Labs, and the open source go-ipfs code, as well as some familiarity of the team with that code, it is realistic to build something like that relatively quickly (workload estimate: 80 hours).

- *Vbeam routing table*: let's build this with a distributed hash table based on the Ethereum smart contracts; while there are probability other good alternatives to tackle the same problem without overengineering this, this take will be more aligned for this given hackathon;  we assume that the ISPs who control a given range of IP addresses have a key that makes them verifiably publish the routing recommendation for a given peerlist's data object  (workload estimate: TBD).
