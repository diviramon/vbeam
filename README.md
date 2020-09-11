# Voronoibeam

The decentralized content delivery based on IPFS for internet service providers (ISPs).

[HackFS Showcase Link](https://hack.ethglobal.co/showcase/voronoibeam-recmYfuu73aADOldW)

## Motivation

At its core,  TCP/IP works like a telephone network: it enables a point-to-point conversation between two network nodes. This works amazingly for many things but is not a good fit for other things. For content delivery, IPFS offers an alternative model that is unambigously technically superior: content-addressable networking. 

Currently, practical content delivery relies on the propreitary vendor-lock CDN model, where content publishers rely on a properietary "edge infrastructure" owned by CDN companies and installed at Internet Service Providers' (ISPs) premises to make their content globally available.

vbeam's design intends to be the best of these two architectures. 

vbeam is a fully decentralized open protocol based on and compatible with IPFS. Thus, it unlocks open, trustless and content-addressable networking. 

At the same time, vbeam has the best of the traditional CDN model: the publisher retaining complete ownership and control of their internet resource and their customers' experience. With vbeam, you are not at the mercy of a fancy algorithm one  need a masters degree to understand. Publishing content, purging and changing their content at any time is done with just some straightforward tweaking of some JSONs and the way your HTTP app behaves, something web developers are comfortable with. 



## Overview

- *Vbeam pinpub*: a publisher announces to the world the list of data object ids (the pinlist) and regularly publishes changes or updates to it; the ISPs that are interested in keeping track of the pinlists for a given publisher are subscribing to such announcements; we can see that generally this sounds as a pubsub model.

- *Vbeam fwd*: the publisher can keep track of the routing tables for their peerlists for all the ISPs and redirect the customers's requests to the gateway recommended by their ISP.

- *Vbeam gateway*: a server to be run within the ISP infrastructure  that is largely a regular IPFS node. The ISP can subscribe their gateway to the publisher's pinlists and which adds the data object ids listed there to its own pinlist; the server also serves the acquired data objects over the HTTP gateway so that the ISP's customers can be served these data objects over HTTP.

- *vbeam.js*: is the embeddable javascript module that the publisher includes to their webpage that carries consumer-side logic needed for vbeam. vbeam.js takes the CIDs of the content distributed over vbeam as an input and sends the request to the pubsub's forwarder which tells it which HTTP gateway to use in order to fetch the content in question;     


- *Vbeam routing table*: the voronoibeam gateways go up and down all the time, subscribe and unsubscribe to pinlists and so on.  The ISPs are the ones who know best which of the gateways the traffic from their customers needs to redirected to, and they publish these rerouting recommendations within the decentralized routing table system.


## Specifications

- *Vbeam pinpub*: we can start with something RSS-style in spirit; that is, serving the pinlists as https://www.google.com/search?q=ip+routing+table&rls=com.microsoft:en-GB:{referrer:source?}&ie=UTF-8&oe=UTF-8&sourceid=ie7&rlz=1I7SKPB_ruJSONs over HTTP and stuff; maybe a magic word DNS record that publishers can add to signify a given (workload estimate: 20 hours).

- *Vbeam gateway*:  thanks to libp2p by the Protocol Labs, and the open source go-ipfs code, as well as some familiarity of the team with that code, it is realistic to build something like that relatively quickly (workload estimate: 80 hours).

- *Vbeam routing table*: let's build this with a distributed hash table based on the Ethereum smart contracts; while there are probability other good alternatives to tackle the same problem without overengineering this, this take will be more aligned for this given hackathon;  we assume that the ISPs who control a given range of IP addresses have a key that makes them verifiably publish the routing recommendation for a given peerlist's data object  (workload estimate: TBD).


## Installation

This thing is a work a progress.

Create a repo for each service with: 

```bash
IPFS_PATH=./env/pinpub/ipfs-repo ipfs init 
IPFS_PATH=./env/gateway/ipfs-repo ipfs init 
```
