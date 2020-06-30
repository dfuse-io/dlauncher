import React, { createContext, useContext, useState } from "react";
const windowTS = window as any

export interface ConfigContextType {
  title: string
  blockExplorerName: string
}

export const ConfigContext = createContext<ConfigContextType>(null!)


export default function ConfigProvider(props: React.PropsWithChildren<{}>) {
  const title = (windowTS.TopLevelConfig ? windowTS.TopLevelConfig.title  : "dfuse dashboard")
  const blockExplorerName = (windowTS.TopLevelConfig ? windowTS.TopLevelConfig.blockExplorerName  : "explorer")

  return (
    <ConfigContext.Provider value={{title, blockExplorerName}}>
      {props.children}
    </ConfigContext.Provider>
  )
}

export function useConfig() {
  const context = useContext(ConfigContext)
  if(!context) {
    throw new Error("usConfig must be used within the ConfigProvider")
  }
  return context
}