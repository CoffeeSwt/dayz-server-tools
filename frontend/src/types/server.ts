export interface Server {
    id?: number // 服务器ID
    name?: string // 服务器名称
    ip?: string // 服务器IP
    port?: number // 服务器端口
    password?: string // 服务器密码
    status?: string // 服务器状态
    newServer?: boolean // 是否为新服务器
}
