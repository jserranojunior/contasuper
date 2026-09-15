import ClassUseApiConnect from "../../../helpers/http/useApiConnect";
const ApiConnect = new ClassUseApiConnect();

export function useTransaction() {

  // Função para solicitar Depósito
  async function createDeposit(data: Record<string, unknown>) {
    const urlApi = "/financial/deposit";
    return await ApiConnect.post(urlApi, data)
      .then((res: any) => handleResponse(res))
      .catch((res: any) => handleError(res));
  }

  // Função para solicitar Saque (futura)
  async function createWithdraw(data: Record<string, unknown>) {
    const urlApi = "/financial/withdraw";
    return await ApiConnect.post(urlApi, data)
      .then((res: any) => handleResponse(res))
      .catch((res: any) => handleError(res));
  }

  // Função para buscar o extrato/transações do usuário
  async function getExtrato() {
    const urlApi = "/financial/extrato";
    return await ApiConnect.get(urlApi)
      .then((res: any) => handleResponse(res))
      .catch((res: any) => handleError(res));
  }

  // Funções auxiliares para padronizar o tratamento:
  function handleResponse(res: any) {
    if (res && res.data) {
      return res.data;
    } else {
      if (res && res.response && res.response.data && res.response.data.message) {
        return { error: res.response.data.message };
      } else {
        return { error: "Erro ao processar a transação, contate o administrador" };
      }
    }
  }

  function handleError(res: any) {
    console.log("Erro de requisição", res);
    return { error: "Erro de conexão com o servidor" };
  }

  return { 
    createDeposit, 
    createWithdraw,
    getExtrato
  };
}