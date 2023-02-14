import React, { MouseEventHandler } from 'react'
import styled from 'styled-components'
import Anchor from './Anchor'

const Container = styled.div`
  display: flex;
  gap: calc(var(--spacing-v1) * 1);
`

type PaginatorProps = {
  enabled?: boolean
  page: number
  pages: number
  onChange: (page: number) => void
}

const Paginator: React.FC<PaginatorProps> = ({
  enabled = true,
  pages,
  page,
  onChange,
}) => {
  const onClick: (n: number) => MouseEventHandler<any> = (page) => {
    return (evt) => {
      evt.preventDefault()

      if (page < 1) return
      if (page > pages) return

      if (onChange) onChange(page)
    }
  }
  return (
    <Container>
      {enabled || pages > 5 ? (
        <Anchor href="" onClick={onClick(1)}>
          <span className="material-symbols-outlined">first_page</span>
        </Anchor>
      ) : null}
      <Anchor href="" onClick={onClick(page - 1)}>
        <span className="material-symbols-outlined">navigate_before</span>
      </Anchor>
      <Anchor href="" onClick={onClick(page + 1)}>
        <span className="material-symbols-outlined">navigate_next</span>
      </Anchor>
      {enabled || pages > 5 ? (
        <Anchor href="" onClick={onClick(pages)}>
          <span className="material-symbols-outlined">last_page</span>
        </Anchor>
      ) : null}
    </Container>
  )
}

export default Paginator
