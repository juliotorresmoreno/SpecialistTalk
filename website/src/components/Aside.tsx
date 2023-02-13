import React from 'react'
import styled from 'styled-components'
import Contacts from './Social/Contacts'

const Container = styled.aside`
  border-left: 1px solid var(--bs-gray-300);
  display: flex;
  flex-direction: column;

  @media (max-width: 767px) {
    display: none;
  }
`

const Aside = () => {
  return (
    <Container>
      <Contacts />
    </Container>
  )
}

export default Aside
