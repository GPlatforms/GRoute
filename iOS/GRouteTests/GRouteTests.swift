//
//  GRouteTests.swift
//  GRouteTests
//
//  Created by weibo on 17/3/16.
//  Copyright © 2017年 lez. All rights reserved.
//

import XCTest
@testable import GRoute

class GRouteTests: XCTestCase {
    
    override func setUp() {
        super.setUp()
        // Put setup code here. This method is called before the invocation of each test method in the class.
    }
    
    override func tearDown() {
        // Put teardown code here. This method is called after the invocation of each test method in the class.
        super.tearDown()
    }
    
    func testPredicate() {
        let gRoute = GRouteManager.sharedInstance
        let time = NSDate().timeIntervalSince1970
        print(time)
        gRoute.getConfig(app_id: "11235", time: time, sign: "", urls: ["http://wenzb.com/api/v1/app/config/dns_info"]) { 
            print(gRoute.getBaseUrl())
            print(gRoute.originDict)
        }
        // This is an example of a functional test case.
        // Use XCTAssert and related functions to verify your tests produce the correct results.
    }
    
    func testDownLoad() {
        
    }
    
    func testMatch() {
        let gRoute = GRouteManager.sharedInstance
    }
    
    func testPerformanceExample() {
        // This is an example of a performance test case.
        self.measure {
            // Put the code you want to measure the time of here.
        }
    }
    
}
