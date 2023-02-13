import React from 'react'
import styled from 'styled-components'
import Contacts from './Contacts'
import Provider from './Provider'

const height = 'calc(var(--spacing-v1) * 3.5)'
const width = '200px'

const Container = styled.div`
  display: flex;
  position: fixed;
  height: ${height};
  line-height: ${height};
  user-select: none;
  background-color: var(--bs-teal);

  transform: translateY(calc(100vh - ${height}))
    translateX(calc((100vw - 540px) / 2));
  width: 540px;

  @media (max-height: 400px) {
    display: none;
  }

  @media (min-width: 768px) {
    transform: translateY(calc(100vh - ${height}))
      translateX(calc((100vw - 720px) / 2 + 520px));
    width: ${width};
  }

  @media (min-width: 992px) {
    transform: translateY(calc(100vh - ${height}))
      translateX(calc((100vw - 960px) / 2 + 760px));
    width: ${width};
  }

  @media (min-width: 1200px) {
    transform: translateY(calc(100vh - ${height}))
      translateX(calc((100vw - 1140px) / 2 + 940px));
    width: ${width};
  }

  @media (min-width: 1400px) {
    transform: translateY(calc(100vh - ${height}))
      translateX(calc((100vw - 1320px) / 2 + 1120px));
    width: ${width};
  }
`

const Social: React.FC = ({}) => {
  return (
    <Provider>
      <Container>
        <Contacts />
      </Container>
    </Provider>
  )
}

export default Social
