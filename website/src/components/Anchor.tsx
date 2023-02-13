import styled from 'styled-components'

const notAction: React.MouseEventHandler<HTMLAnchorElement> = (evt) =>
  evt.preventDefault()

const Anchor = styled.a.attrs((props) => ({
  onClick: notAction,
  ...props,
}))`
  color: var(--bs-gray-700);
  cursor: pointer;
  user-select: none;
`

export default Anchor
