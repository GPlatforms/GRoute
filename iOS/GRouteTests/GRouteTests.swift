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
        let gRoute = GRoute.sharedInstance
        XCTAssert(gRoute.textMatch(text: "asdf!", pattern: ".*"))
        
        XCTAssertEqual((gRoute.textMatch(text: "asdf", pattern: "x")), false)
        
        // This is an example of a functional test case.
        // Use XCTAssert and related functions to verify your tests produce the correct results.
    }
    
    func testDownLoad() {
        
    }
    
    func testMatch() {
        let gRoute = GRoute.sharedInstance
        
        gRoute.routeConfig = [Rule(newReg: "fa", newURL: "http://www.baidu.com"),
                              Rule(newReg: ".*", newURL: "http://www.taobao.com")]
        let urlA = gRoute.match(functionName: "fa")
        let urlB = gRoute.match(functionName: "fb")
        print(urlA)
        print(urlB)
        XCTAssertEqual(urlA, "http://www.baidu.com")
        XCTAssertEqual(urlB, "http://www.taobao.com")
    }
    
    func testPerformanceExample() {
        // This is an example of a performance test case.
        self.measure {
            // Put the code you want to measure the time of here.
        }
    }
    
}
