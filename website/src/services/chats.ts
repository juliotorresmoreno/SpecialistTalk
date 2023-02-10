import React from 'react'
import config from '../config'
import { useGetData } from './api'

export function useGetChat(id: string) {
  const url = config.baseUrl + '/chats/'
  const { data, error, isLoading } = useGetData(url + id)

  return { isLoading, error, data }
}
