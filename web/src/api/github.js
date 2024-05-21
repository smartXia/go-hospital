import axios from 'axios'

const service = axios.create()

export function Commits(page) {
  return service({
    url: "https://github.com/jumpserver/jumpserver/commits?page=" + page,
    method: "get",
  });
}

export function Members() {
  return service({
    url: "https://github.com/jumpserver/jumpserver/members",
    method: "get",
  });
}
