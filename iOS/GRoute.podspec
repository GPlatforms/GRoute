Pod::Spec.new do |s|
    s.name         = 'GRoute'
    s.version      = '1.0'
    s.summary      = 'Router config'
    s.homepage     = 'https://github.com/GPlatforms/GRoute'
    s.authors      = {'weibo' => 'weibo3721@126.com'}
    s.ios.deployment_target = '8.0'
    s.source       = { :git => 'https://github.com/GPlatforms/GRoute.git'}
    s.source_files = 'iOS/GRoute/*.swift'
    s.dependency 'Alamofire'
end
