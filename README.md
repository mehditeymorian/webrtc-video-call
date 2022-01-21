# webrtc-video-call
This project is a video calling application using WebRTC technology. It contains minimum requirements to provide a group video call.
[WebRTC](https://webrtc.org/) is a free and open-source project providing web browsers and mobile applications with real-time communication
via application programming interfaces.


# Project Sections.
WebRTC is direct peer-to-peer connection which is done using WebRTC API in browsers and mobile applications.
But for starting a group call, it requires a mechnaism to share peers address and connectivity information with each other. This mechanism is called Signaling Server.
It keeps an open connection with each peer to exchange peers' address, connectivity info, and room events.       

This project uses Websockets to keep an open connection with peers. Also there are HTTP endpoints for room management.        
Whenever a peer joins in a room, a join event send to other peers in the room, then other peers create offers and sent it to new peer; 
The new peer answer the offer and a connection takes place(it is a lot more complicated than this, please read official documents and take a look at code to uderstand more). 

# How to run
The frontend is a collection of html files that are served as static files. build and run the project and you are good to go.
- Create or Join Room -> `ADDRESS/`
- Room                -> `ADDRESS/room?roomId=[ROOM-ID]&name=[PEER-NAME]`

