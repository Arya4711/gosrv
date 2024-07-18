export default async function getUsers() {
  const response = await fetch("http://localhost:5050/users", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });

  return response.json();
}
