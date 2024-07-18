import { User } from "./models";

export default async function getUsers() {
  const response = await fetch("http://localhost:5050/users", {
    method: "GET",
  });

  return (await response.json()) as User[];
}
