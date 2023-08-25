// 检查用户名是否合法
const checkToUserName = (username: string) =>{
    // 数字、字母、下划线，一种及以上，长度为 4-20
    const reg = /^\w{2,20}$/g  
    return reg.test(username)
}

const checkToPassword = (password: string) => {
    // 数字、字母、下划线至少包含两种，长度为 6-20
    //const reg = /(?!^\d+$)(?!^[A-Za-z]+$)(?!^_+$)^\w{8,20}$/g;
    const reg = /^\w{4,20}$/g  
    return reg.test(password);
  };
  

export default  {
    checkToUserName,
    checkToPassword,
  };