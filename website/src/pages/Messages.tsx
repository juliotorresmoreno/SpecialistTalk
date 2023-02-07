import React from 'react'
import Header from '../components/Header'
import withChat from '../hoc/withChat'

const _MessagesPage: React.FC = () => {
  const header = {
    title: 'Messages',
    description: 'programa de super poderes',
  }
  return (
    <>
      <Header {...header} />
      <main>dfasd</main>
    </>
  )
}

const MessagesPage = withChat(_MessagesPage)

export default MessagesPage
