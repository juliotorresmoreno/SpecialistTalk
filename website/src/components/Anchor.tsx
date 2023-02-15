import styled from 'styled-components'

const notAction: React.MouseEventHandler<HTMLAnchorElement> = (evt) =>
  evt.preventDefault()

const Anchor = styled.a.attrs((props) => ({
  ...props,
  onClick: props.onClick ?? notAction,
}))`
  color: var(--bs-gray-700);
  cursor: pointer;
  user-select: none;
`

export default Anchor
