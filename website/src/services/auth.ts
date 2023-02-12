import React from 'react'
import config from '../config'
import { useGetData } from './api'

export function useGetSession() {
  const url = config.baseUrl + '/auth/session'
  const { error, isLoading, get } = useGetData(url)

  return { isLoading, error, get }
}
