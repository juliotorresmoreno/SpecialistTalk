import moment from 'moment'
import React from 'react'
import styled from 'styled-components'
import { Message as _Message } from '../../models/message'
import { useAppSelector } from '../../store/hooks'

type MessageProps = {
  data: _Message
}

type ContainerProps = {
  isMe: boolean
}
const Container = styled.div<ContainerProps>`
  text-align: ${(props) => (props.isMe ? 'end' : 'start')};
`

const DateContainer = styled.div`
  font-size: calc(var(--font-size) * 0.7);
  color: var(--bs-gray-700);
`

const ContainerName = styled.span`
  font-weight: bold;
`

const Message: React.FC<MessageProps> = ({ data }) => {
  const user = useAppSelector((state) => state.auth.session?.user)
  const date = moment(data.created_at).format('MMMM Do YYYY, h:mm:ss a')
  const isMe = user?.id === data.from
  const message = isMe ? (
    <>
      {data.message}: <ContainerName>{data.from_name}</ContainerName>
    </>
  ) : (
    <>
      <ContainerName>{data.from_name}</ContainerName>: {data.message}
    </>
  )
  return (
    <Container isMe={isMe}>
      <div>{message}</div>
      <DateContainer>{date}</DateContainer>
    </Container>
  )
}

export default Message
