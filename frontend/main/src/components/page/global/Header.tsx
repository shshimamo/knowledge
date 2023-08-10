import { useCurrentUserState } from '@/globalStates/currentUserState'

export const Header = () => {
  const currentUser = useCurrentUserState();

  const handleSignout = async () => {
    // TODO: SignoutUsecase
  };

  return (
    <header className="bg-blue-500 text-white px-6 py-4 flex justify-between items-center">
      <h1 className="font-bold text-2xl">Your App Name</h1>
      {currentUser && (
        <div className="flex items-center">
          <p className="mr-4">Welcome, {currentUser.name}!</p>
          <button className="bg-red-600 hover:bg-red-700 text-white font-bold py-2 px-4 rounded" onClick={handleSignout}>
            Logout
          </button>
        </div>
      )}
    </header>
  );
}