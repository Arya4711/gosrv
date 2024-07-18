import getUsers from "./_lib/getUsers";
import { User } from "./_lib/models";

export default async function Home() {
  const response = await fetch("http://localhost:5050/users", {
    method: "GET",
  });

  console.log(response);
  console.log(await response.json());

  return <div></div>;
}
