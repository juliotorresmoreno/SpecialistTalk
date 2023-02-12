import React from 'react'
import styled from 'styled-components'
import Contacts from './Social/Contacts'

const Container = styled.aside`
  border-left: 1px solid var(--bs-gray-300);
  display: flex;
  flex-direction: column;
`

const Ads = styled.div`
  flex: 1;
`

const Aside = () => {
  return (
    <Container>
      <Ads></Ads>

      <Contacts />
    </Container>
  )
}

export default Aside
