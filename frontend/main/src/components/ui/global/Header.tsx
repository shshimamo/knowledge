import { useCurrentUserState } from '@/globalStates/currentUserState'

export const Header = () => {
  const currentUser = useCurrentUserState()

  const handleSignout = () => {
    // TODO: SignoutUsecase
  }

  return (
    <header className='flex items-center justify-between bg-blue-500 px-6 py-4 text-white'>
      <h1 className='text-2xl font-bold'>Your App Name</h1>
      {currentUser.id && (
        <div className='flex items-center'>
          <p className='mr-4'>Welcome, {currentUser.name}!</p>
          <button
            className='rounded bg-red-600 px-4 py-2 font-bold text-white hover:bg-red-700'
            onClick={handleSignout}
          >
            Logout
          </button>
        </div>
      )}
    </header>
  )
}
