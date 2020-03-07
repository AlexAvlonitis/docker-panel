export default [
  {
    _name: 'CSidebarNav',
    _children: [
      {
        _name: 'CSidebarNavTitle',
        _children: ['Panel']
      },
      {
        _name: 'CSidebarNavItem',
        name: 'Images',
        to: '/images',
        icon: 'cilLayers'
      },
      {
        _name: 'CSidebarNavItem',
        name: 'Containers',
        to: '/containers',
        icon: 'cilPuzzle'
      },
      {
        _name: 'CSidebarNavItem',
        name: 'Networks',
        to: '/networks',
        icon: 'cilGrid'
      }
    ]
  }
]
