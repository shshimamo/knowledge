import { SigninReqData, SigninResData, SignupReqData, SignupResData } from '@/api/auth/apiClient'
import { SigninSeed, SignupSeed, Token } from '@/components/model/auth/type'

export const convertSignupSeedToReqData = ({ email, password }: SignupSeed): SignupReqData => {
  return { email, password }
}

export const convertSignupResDataToToken = (resData: SignupResData): Token => {
  return resData.token
}

export const convertSigninSeedToReqData = ({ email, password }: SigninSeed): SigninReqData => {
  return { email, password }
}

export const convertSigninResDataToToken = (resData: SigninResData): Token => {
  return resData.token
}
