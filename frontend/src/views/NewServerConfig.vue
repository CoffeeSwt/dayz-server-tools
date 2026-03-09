<template>
    <div>1111111</div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
const serverConfig = reactive({
    serverName: '', //服务器名称
    mpmissionsPath: '', //服务器mpmissions路径
    mpmissionsCfgPath: '', //服务器mpmissionsCfg路径
    port: '', //服务器端口
})

const officialMapLinks = (mapNameEng:string)=>{
    const map = new Map()
    map.set('chernarusplus', {
        mpmissionsPath:'/mpmissions/dayzOffline.chernarusplus',
        mpmissionsCfgPath:'/mpmissionsCfg/chernarusplus',
    })
    map.set('dayz_epoch', 'https://dayz-epoch.fandom.com/wiki/Map')
    return map.get(mapNameEng) || ''
}
</script>