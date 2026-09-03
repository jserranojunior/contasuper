import { reactive, toRefs, ref } from "vue";
import { useHttpAuth } from "./useHttpAuth";
import { store } from "./storeAuth";
const { login } = useHttpAuth();

export function useAuth() {
  function getOla() {
    return store.ola;
  }
  function setOla() {
    store.ola = "ada";
  }
  function setStateToken(token: string) {
    store.auth.token = token;
  }
  function setStateAuthError(erro: string) {
    store.auth.erro = erro;
  }
  function setLocalStorageToken(token: string) {
    if (checkOnBrowser()) {
      localStorage.setItem("token", token);
    }
  }
  function checkTokenIsEmpty(token: string) {
    if (token !== "") {
      setStateAuthError("");
    } else {
      setStateAuthError("Erro ao fazer login");
    }
  }
  function checkLocalstorageToken() {
    if (checkOnBrowser()) {
      if (
        localStorage.getItem("token") != "null" ||
        localStorage.getItem("token") != "undefined" ||
        localStorage.getItem("token") != null ||
        localStorage.getItem("token") != undefined
      ) {
        return true;
      } else {
        return false;
      }
    } else {
      return false;
    }
  }
  function checkOnBrowser() {
    if (typeof window !== "undefined") {
      return true;
    } else {
      return false;
    }
  }
  function setToken(token: string) {
    setLocalStorageToken(token);
    setStateToken(token);
  }
  function Logout() {
    setStateAuthError("");
    setToken("");
    /* redirectPageTo("financeiro") */
  }
  function checkStateToken() {
    if (
      !store.auth ||
      !store.auth.token ||
      store.auth.token == "" ||
      store.auth.token == undefined ||
      store.auth.token == "undefined" ||
      store.auth.token == "null" ||
      store.auth.token == null ||
      store.auth.token.length == 0
    ) {
      return false;
    } else {
      return true;
    }
  }
  function setTokenEqualStorageState() {
    if (
      checkOnBrowser() &&
      checkLocalstorageToken() &&
      localStorage.getItem("token") !== store.auth.token
    ) {
      setToken(String(localStorage.getItem("token")));
    }
  }
  async function isLogged() {
    setTokenEqualStorageState();
    const checked = checkStateToken();
    return checked;
  }
  function checkFieldsIsValid() {
    if (store.fields && store.fields.email > "" && store.fields.password > "") {
      return true;
    } else {
      setStateAuthError("Campos Vazios");
      setToken("");
      return false;
    }
  }

  function setStateFields(dataFields: { email: string; password: string }) {
    store.fields.email = dataFields.email;
    store.fields.password = dataFields.password;
  }

  async function Login() {
    setStateAuthError("");
    if (checkFieldsIsValid()) {
      return await login(store.fields).then((res) => {
        // Verifique se a API retornou o token e os dados do usuário
        if (res && res.data && res.data.token) {
          setToken(res.data.token);
          
          // SALVA OS DADOS DO USUÁRIO NO STORE (Se o backend enviar em res.data.user)
          if (res.data.user) {
            store.auth.user = res.data.user;
            
            // Opcional: Se quiser salvar no localStorage para persistir ao atualizar a página
            if (checkOnBrowser()) {
              localStorage.setItem("user_name", res.data.user.name);
              localStorage.setItem("user_cellphone", res.data.user.cellphone);
            }
          }
          
          return true;
        } else {
          // ... (seus tratamentos de erro existentes)
          return false;
        }
      });
    } else {
      return false;
    }
  }

  return {
    ...toRefs(store),
    setOla,
    checkOnBrowser,
    setToken,
    checkStateToken,
    checkFieldsIsValid,
    setStateFields,
    setLocalStorageToken,
    setStateToken,
    getOla,
    checkLocalstorageToken,
    setStateAuthError,
    checkTokenIsEmpty,
    Logout,
    Login,
    isLogged,
  };
}
