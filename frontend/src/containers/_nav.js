export default [
  {
    _name: 'CSidebarNav',
    _children: [
      {
        _name: 'CSidebarNavTitle',
        _children: ['Docker Views']
      },
      {
        _name: 'CSidebarNavItem',
        name: 'Images',
        to: '/images',
        icon: 'cil-drop'
      }
    ]
  }
]
