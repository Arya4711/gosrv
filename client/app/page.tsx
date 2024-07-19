import { useRouter } from "next/navigation";
import addUser from "./_lib/addUser";
import getUsers from "./_lib/getUsers";
import { User } from "./_lib/models";

export default async function Home() {
  const users = await getUsers();

  return (
    <>
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
        <label htmlFor="username" className="block">
          Username
        </label>
        <input
          id="username"
          name="username"
          className="dark:text-black"
          required
        />
        <label htmlFor="password" className="block">
          Password
        </label>
        <input
          id="password"
          name="password"
          className="dark:text-black"
          required
        />
        <button type="submit" className="block border">
          Submit
        </button>
      </form>
    </>
  );
}
