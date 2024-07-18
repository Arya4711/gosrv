import getUsers from "./_lib/getUsers";
import { User } from "./_lib/models";

export default async function Home() {
  const users = await getUsers();

  return (
    <>
      <ul>
        {users.map((user: User) => (
          <li key={user.id}>
            id: {user.id} username: {user.username} password: {user.password}
          </li>
        ))}
      </ul>
    </>
  );
}
