import { SigninAPIReqParams, SigninAPIResponse, SignupAPIReqParams, SignupAPIResponse } from '@/api/auth/apiClient'
import { SigninSeed, SignupSeed, Token } from '@/components/model/auth/type'

export const convertSignupSeedToReqParams = ({ email, password }: SignupSeed): SignupAPIReqParams => {
  return { email, password }
}

export const convertSignupResponseToToken = (resData: SignupAPIResponse): Token => {
  return resData.token
}

export const convertSigninSeedToReqParams = ({ email, password }: SigninSeed): SigninAPIReqParams => {
  return { email, password }
}

export const convertSigninResponseToToken = (resData: SigninAPIResponse): Token => {
  return resData.token
}
