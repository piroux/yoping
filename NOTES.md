
## FLOW

FLOW Send:
- User-From sends ping from Mobile-App (Mobile-App)
- Mobile-App of User-From sends Ping for User-To to Backend-App (Mobile-App)
- Backend-App receives Ping, stores it in DB and notify Mobile-App of User-To (Backend-App: receivePing)
  - PingExRequest
  - PingExResponse
- User-To receives notification and open Mobile-App (Mobile-App)
- Mobile-App of User-To fetches Ping info from Backend-App (Mobile-App)
- Backend-App returns with info for Ping & deletes Ping between User-From and User-To in DB (Backend-App: sendPing)
  - PingInRequest
  - PingInResponse
