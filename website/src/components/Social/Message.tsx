import moment from 'moment'
import React from 'react'
import styled from 'styled-components'
import { Message as _Message } from '../../models/message'
import { useAppSelector } from '../../store/hooks'

type MessageProps = {
  data: _Message
}

type WithMe = {
  isMe: boolean
}
const Container = styled.div<WithMe>`
  display: flex;
  justify-content: ${(props) => (props.isMe ? 'start' : 'end')};
  margin-bottom: var(--spacing-v1);
`

const DateContainer = styled.div`
  font-size: calc(var(--font-size) * 0.7);
  color: var(--bs-gray-700);
`

const ContainerName = styled.span`
  font-weight: bold;
`

const Bubble = styled.div<WithMe>`
  background-color: var(--bs-gray-200);
  flex: 0.8;
  border-radius: 5px;
  padding: var(--spacing-v1);
  text-align: ${(props) => (props.isMe ? 'start' : 'end')};
`

const Message: React.FC<MessageProps> = ({ data }) => {
  const user = useAppSelector((state) => state.auth.session?.user)
  const date = moment(data.created_at).format('MMMM Do YYYY, h:mm:ss a')
  const isMe = user?.id === data.from
  const message = isMe ? (
    <>
      <ContainerName>{data.from_name}</ContainerName>: {data.message}
    </>
  ) : (
    <>
      {data.message}: <ContainerName>{data.from_name}</ContainerName>
    </>
  )
  return (
    <Container isMe={isMe}>
      <Bubble isMe={isMe}>
        <div>{message}</div>
        <DateContainer>{date}</DateContainer>
      </Bubble>
    </Container>
  )
}

export default Message
