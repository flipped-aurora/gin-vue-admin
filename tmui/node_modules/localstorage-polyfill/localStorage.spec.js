/* global localStorage */
import test from 'ava'
import './localStorage'

test('methods should work', (t) => {
  // can't make assuptions about key positioning
  localStorage.setItem('a', 1)
  t.is(localStorage.key(0), 'a')

  localStorage.setItem('b', '2')
  t.is(localStorage.getItem('a'), '1')
  t.is(localStorage.getItem('b'), '2')
  t.is(localStorage.length, 2)

  t.is(localStorage['c'], undefined)
  t.is(localStorage.getItem('c'), null)

  localStorage.setItem('c')
  t.is(localStorage.getItem('c'), 'undefined')
  t.is(localStorage.length, 3)

  localStorage.removeItem('c')
  t.is(localStorage.getItem('c'), null)
  t.is(localStorage.length, 2)

  localStorage.clear()
  t.is(localStorage.getItem('a'), null)
  t.is(localStorage.getItem('b'), null)
  t.is(localStorage.length, 0)
})

test('error throwing', (t) => {
  t.throws(() => {
    localStorage.key()
  }, /Failed to execute 'key' on 'Storage': 1 argument required, but only 0 present\./)
})

test('proxy should work', (t) => {
  t.is(localStorage.length, 0)

  localStorage.a = {}
  t.is(localStorage.a, '[object Object]')
  localStorage.c = 1
  const obj = {}
  localStorage[obj] = 'key gets stringified'
  t.is(localStorage['[object Object]'], 'key gets stringified')
  t.is(localStorage.c, '1')
  t.is(localStorage.length, 3)
  localStorage.length = 0
  t.is(localStorage.length, 3)
  localStorage.key = 'only an ass**** would do this'
  t.is(localStorage.length, 3)
  t.is(localStorage.key, 'only an ass**** would do this')
})
