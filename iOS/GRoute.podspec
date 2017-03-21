Pod::Spec.new do |s|
    s.name         = 'GRoute'
    s.version      = '1.0'
    s.summary      = 'Router config'
    s.homepage     = 'https://github.com/GPlatforms/GRoute'
    s.license      = 'MIT'
    s.authors      = {'weibo' => 'weibo3721@126.com'}
    s.platform     = :ios, '8.0'
    s.source = { :git => 'https://github.com/Alamofire/Alamofire.git', :tag => s.version }
    s.source_files = 'iOS/GRoute/*.swift'
end