import { addUser, deleteUser, getUsers } from "./_lib/users";
import { User } from "./_lib/models";

export default async function Home() {
  const users = await getUsers();

  return (
    <div className="m-5 space-y-5">
      <ul>
        {users
          ? users.map((user: User) => (
              <li key={user.id}>
                id: {user.id} username: {user.username} password:{" "}
                {user.password}
              </li>
            ))
          : "no users found"}
      </ul>
      <form action={addUser}>
        <h1 className="text-xl">add user</h1>
        <label htmlFor="username" className="block">
          username
        </label>
        <input
          id="username"
          name="username"
          className="dark:text-black"
          required
        />
        <label htmlFor="password" className="block">
          password
        </label>
        <input
          id="password"
          name="password"
          className="dark:text-black"
          required
        />
        <button type="submit" className="block border">
          submit
        </button>
      </form>
      <form action={deleteUser}>
        <h1 className="text-xl">delete user</h1>
        <label htmlFor="username" className="block">
          username
        </label>
        <input
          id="username"
          name="username"
          className="dark:text-black"
          required
        />
        <button type="submit" className="block border">
          submit
        </button>
      </form>
    </div>
  );
}
