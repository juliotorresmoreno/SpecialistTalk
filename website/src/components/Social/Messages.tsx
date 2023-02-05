import React from 'react'
import styled from 'styled-components'
import Header from '../Header'

const Container = styled.div`
  background-color: white;
`
export const Messages = () => {
  const header = {
    title: 'Messages',
    description: 'programa de super poderes',
  }
  return (
    <Container>
      <Header {...header} />
    </Container>
  )
}
