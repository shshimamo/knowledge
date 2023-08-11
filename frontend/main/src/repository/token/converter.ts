import { SetTokenAPIReqParams } from '@/api/main/apiClient'
import { Token } from '@/components/model/auth/type'

export const convertTokenToReqParams = (token: Token): SetTokenAPIReqParams => {
  return { token }
}