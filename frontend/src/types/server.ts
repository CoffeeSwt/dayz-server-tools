export interface Server {
    id?: number // 服务器ID
    name?: string // 服务器名称
    port?: number // 服务器端口
    map?: string // 服务器地图
    newServer?: boolean // 是否为新服务器
}

const officialMapLinks = (mapNameEng:string)=>{
    const map = new Map()
    map.set('chernarusplus', {
        mpmissionsPath:'/mpmissions/dayzOffline.chernarusplus',
        mpmissionsCfgPath:'/serverCfg/chernarusplus.cfg',
    })
    map.set('dayz_epoch', 'https://dayz-epoch.fandom.com/wiki/Map')
    return map.get(mapNameEng) || ''
}
