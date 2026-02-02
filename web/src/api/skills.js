import service from '@/utils/request'

export const getSkillTools = () => {
  return service({
    url: '/skills/getTools',
    method: 'get'
  })
}

export const getSkillList = (data) => {
  return service({
    url: '/skills/getSkillList',
    method: 'post',
    data
  })
}

export const getSkillDetail = (data) => {
  return service({
    url: '/skills/getSkillDetail',
    method: 'post',
    data
  })
}

export const saveSkill = (data) => {
  return service({
    url: '/skills/saveSkill',
    method: 'post',
    data
  })
}

export const createSkillScript = (data) => {
  return service({
    url: '/skills/createScript',
    method: 'post',
    data
  })
}

export const getSkillScript = (data) => {
  return service({
    url: '/skills/getScript',
    method: 'post',
    data
  })
}

export const saveSkillScript = (data) => {
  return service({
    url: '/skills/saveScript',
    method: 'post',
    data
  })
}

export const createSkillResource = (data) => {
  return service({
    url: '/skills/createResource',
    method: 'post',
    data
  })
}

export const getSkillResource = (data) => {
  return service({
    url: '/skills/getResource',
    method: 'post',
    data
  })
}

export const saveSkillResource = (data) => {
  return service({
    url: '/skills/saveResource',
    method: 'post',
    data
  })
}

export const createSkillReference = (data) => {
  return service({
    url: '/skills/createReference',
    method: 'post',
    data
  })
}

export const getSkillReference = (data) => {
  return service({
    url: '/skills/getReference',
    method: 'post',
    data
  })
}

export const saveSkillReference = (data) => {
  return service({
    url: '/skills/saveReference',
    method: 'post',
    data
  })
}

export const createSkillTemplate = (data) => {
  return service({
    url: '/skills/createTemplate',
    method: 'post',
    data
  })
}

export const getSkillTemplate = (data) => {
  return service({
    url: '/skills/getTemplate',
    method: 'post',
    data
  })
}

export const saveSkillTemplate = (data) => {
  return service({
    url: '/skills/saveTemplate',
    method: 'post',
    data
  })
}

export const getGlobalConstraint = (data) => {
  return service({
    url: '/skills/getGlobalConstraint',
    method: 'post',
    data
  })
}

export const saveGlobalConstraint = (data) => {
  return service({
    url: '/skills/saveGlobalConstraint',
    method: 'post',
    data
  })
}
