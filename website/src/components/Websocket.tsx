import React, { useEffect } from 'react'
import config from '../config'
import { wsHandlers } from '../handlers'
import { useAppSelector } from '../store/hooks'

type WebsocketProps = {} & React.PropsWithChildren

const hOpen = wsHandlers.filter((x) => x.event === 'open')
const hMessage = wsHandlers.filter((x) => x.event === 'message')
const hError = wsHandlers.filter((x) => x.event === 'error')
const hClose = wsHandlers.filter((x) => x.event === 'close')

const Websocket: React.FC<WebsocketProps> = ({ children }) => {
  let unmount = false
  let socket: WebSocket | null = null
  const session = useAppSelector((state) => state.auth.session)

  const reconnect = () => {
    if (!session || !session.token) {
      if (socket !== null) {
        socket.close()
      }
      return
    }
    if (socket !== null) return
    socket = new WebSocket(config.wsUrl + '?token=' + session?.token)
    socket.addEventListener('open', function (evt) {
      let _socket = evt.currentTarget as WebSocket
      if (_socket !== socket) _socket.close()

      hOpen.forEach(({ handler }) => handler())
    })

    socket.addEventListener('message', function (evt) {
      let _socket = evt.currentTarget as WebSocket
      if (_socket !== socket) _socket.close()

      const data = JSON.parse(evt.data)

      hMessage.forEach((h) => {
        if ((h as any).type === data.type) {
          h.handler(data)
        }
      })
    })

    socket.addEventListener('error', function (evt) {
      let _socket = evt.currentTarget as WebSocket
      if (_socket !== socket) _socket.close()

      hError.forEach(({ handler }) => handler())
    })

    socket.addEventListener('close', function () {
      if (unmount) return

      socket = null
      setTimeout(() => reconnect(), 3000)

      hClose.forEach(({ handler }) => handler())
    })
  }

  useEffect(() => {
    reconnect()
    return () => {
      unmount = true
      socket?.close()
    }
  })

  return <>{children}</>
}

export default Websocket
