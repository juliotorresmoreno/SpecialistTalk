import styled from 'styled-components'

const Select = styled.select`
  height: calc(var(--spacing-v1) * 3.5);
  border-radius: 0;
  box-shadow: none;
  border: none;
  border-width: 0;
  width: 100%;
  -moz-box-shadow: none;
  -goog-ms-box-shadow: none;
  -webkit-box-shadow: none;
  padding: 0 var(--spacing-v1);
  background-color: var(--bs-gray-200);
  &:hover,
  &:focus,
  &:not(:focus) {
    border-radius: none;
    box-shadow: none;
    border: none;
    border-width: none;
    outline: none;

    -moz-box-shadow: none;
    -goog-ms-box-shadow: none;
    -webkit-box-shadow: none;
  }
`

export default Select
