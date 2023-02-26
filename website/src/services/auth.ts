import React from 'react'
import config from '../config'
import { useGetData, useRemove } from './api'

export function useGetSession() {
  const url = config.baseUrl + '/auth/session'
  const { error, isLoading, get } = useGetData(url)

  return { isLoading, error, get }
}

export function useLogout() {
  const url = config.baseUrl + '/auth/session'
  const { error, isLoading, remove } = useRemove(url)

  return { isLoading, error, apply: () => remove(null) }
}
