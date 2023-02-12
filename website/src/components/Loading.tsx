import React from 'react'
import styled from 'styled-components'

type LoadingProps = {}

const Container = styled.div`
  display: flex;
  flex: 1;
  align-items: center;
  justify-content: center;
`

const Icon = styled.span`
  font-size: calc(var(--spacing-v1) * 8);
`

const Loading: React.FC<LoadingProps> = () => {
  return (
    <Container>
      <Icon className="material-symbols-outlined">hourglass_empty</Icon>
    </Container>
  )
}

export default Loading
