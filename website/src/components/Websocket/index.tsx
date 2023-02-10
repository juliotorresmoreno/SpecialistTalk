import React, { useEffect } from 'react'
import config from '../../config'
import { useAppSelector } from '../../store/hooks'
import Provider from './Provider'

type WebsocketProps = {} & React.PropsWithChildren

const Websocket: React.FC<WebsocketProps> = ({ children }) => {
  const session = useAppSelector((state) => state.auth.session)
  let unmount = false
  let socket: WebSocket | null = null

  const reconnect = () => {
    if (socket !== null || !session?.token) return
    socket = new WebSocket(config.wsUrl + '?token=' + session?.token)
    socket.addEventListener('open', function (evt) {
      let _socket = evt.currentTarget as WebSocket
      if (_socket !== socket) _socket.close()

      console.log('onopen')
    })

    socket.addEventListener('message', function (evt) {
      let _socket = evt.currentTarget as WebSocket
      if (_socket !== socket) _socket.close()

      console.log('message', evt.data)
    })

    socket.addEventListener('error', function (evt) {
      let _socket = evt.currentTarget as WebSocket
      if (_socket !== socket) _socket.close()

      console.log('error')
    })

    socket.addEventListener('close', function (evt) {
      if (unmount) return

      socket = null
      setTimeout(() => reconnect(), 3000)
    })
  }

  useEffect(() => {
    reconnect()
    return () => {
      unmount = true
      socket?.close()
    }
  })

  return <Provider>{children}</Provider>
}

export default Websocket
