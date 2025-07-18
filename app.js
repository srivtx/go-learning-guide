// Application state
let currentScore = 0;
let completedExercises = new Set();
let roadmapProgress = {};

// Exercise data
const exercises = [
    {
        question: "Complete the function signature to add two integers:",
        code: "func add(a, b ____) ____ {\n    return a + b\n}",
        options: ["int, int", "int, string", "string, int"],
        correct: 0,
        explanation: "Both parameters should be 'int' type, and the return type should also be 'int'."
    },
    {
        question: "What is the output of this code?",
        code: "package main\n\nimport \"fmt\"\n\nfunc main() {\n    x := 5\n    if x > 3 {\n        fmt.Println(\"Hello\")\n    } else {\n        fmt.Println(\"World\")\n    }\n}",
        options: ["Hello", "World", "Error"],
        correct: 0,
        explanation: "Since x (5) is greater than 3, the condition is true and 'Hello' is printed."
    },
    {
        question: "Which keyword is used to create a goroutine?",
        code: "// Start a goroutine\n____ myFunction()",
        options: ["go", "async", "thread"],
        correct: 0,
        explanation: "The 'go' keyword is used to start a goroutine in Go."
    }
];

// Wait for DOM to be fully loaded
document.addEventListener('DOMContentLoaded', function() {
    console.log('DOM loaded, initializing app...');
    initializeApp();
});

function initializeApp() {
    console.log('Initializing app...');
    setupTabNavigation();
    setupOSTabNavigation();
    setupCodeExamples();
    setupCopyButtons();
    setupRoadmapProgress();
    setupExercises();
    loadProgress();
    console.log('App initialized successfully');
}

// Tab Navigation - Fixed
function setupTabNavigation() {
    console.log('Setting up tab navigation...');
    const tabBtns = document.querySelectorAll('.tab-btn');
    const tabContents = document.querySelectorAll('.tab-content');
    
    console.log('Found tab buttons:', tabBtns.length);
    console.log('Found tab contents:', tabContents.length);

    tabBtns.forEach((btn, index) => {
        console.log(`Setting up tab button ${index}:`, btn.dataset.tab);
        btn.addEventListener('click', function(e) {
            e.preventDefault();
            const targetTab = this.dataset.tab;
            console.log('Tab clicked:', targetTab);
            
            // Remove active class from all tabs and contents
            tabBtns.forEach(b => b.classList.remove('active'));
            tabContents.forEach(c => c.classList.remove('active'));
            
            // Add active class to clicked tab and corresponding content
            this.classList.add('active');
            const targetContent = document.getElementById(targetTab);
            if (targetContent) {
                targetContent.classList.add('active');
                console.log('Successfully switched to tab:', targetTab);
            } else {
                console.error('Target content not found:', targetTab);
            }
            
            // Save current tab to localStorage
            try {
                localStorage.setItem('currentTab', targetTab);
            } catch (e) {
                console.warn('Could not save to localStorage:', e);
            }
        });
    });
}

// OS Tab Navigation - Fixed
function setupOSTabNavigation() {
    console.log('Setting up OS tab navigation...');
    const osTabs = document.querySelectorAll('.os-tab');
    const osInstructions = document.querySelectorAll('.os-instructions');
    
    console.log('Found OS tabs:', osTabs.length);
    console.log('Found OS instructions:', osInstructions.length);

    osTabs.forEach(tab => {
        tab.addEventListener('click', function(e) {
            e.preventDefault();
            const targetOS = this.dataset.os;
            console.log('OS tab clicked:', targetOS);
            
            // Remove active class from all OS tabs and instructions
            osTabs.forEach(t => t.classList.remove('active'));
            osInstructions.forEach(i => i.classList.remove('active'));
            
            // Add active class to clicked tab and corresponding instructions
            this.classList.add('active');
            const targetInstructions = document.getElementById(targetOS);
            if (targetInstructions) {
                targetInstructions.classList.add('active');
                console.log('Successfully switched to OS:', targetOS);
            } else {
                console.error('Target OS instructions not found:', targetOS);
            }
        });
    });
}

// Code Examples Expandable Sections - Fixed
function setupCodeExamples() {
    console.log('Setting up code examples...');
    const exampleHeaders = document.querySelectorAll('.example-header');
    
    console.log('Found example headers:', exampleHeaders.length);
    
    exampleHeaders.forEach(header => {
        header.addEventListener('click', function(e) {
            e.preventDefault();
            const targetId = this.dataset.toggle;
            const content = document.getElementById(targetId);
            const icon = this.querySelector('.toggle-icon');
            
            console.log('Example header clicked:', targetId);
            
            if (content && icon) {
                // Toggle expanded state
                this.classList.toggle('expanded');
                content.classList.toggle('expanded');
                
                // Update icon
                if (this.classList.contains('expanded')) {
                    icon.textContent = '▲';
                } else {
                    icon.textContent = '▼';
                }
                
                console.log('Toggled example:', targetId, 'expanded:', this.classList.contains('expanded'));
            } else {
                console.error('Content or icon not found for:', targetId);
            }
        });
    });
}

// Copy to Clipboard Functionality - Fixed
function setupCopyButtons() {
    console.log('Setting up copy buttons...');
    const copyBtns = document.querySelectorAll('.copy-btn');
    
    console.log('Found copy buttons:', copyBtns.length);
    
    copyBtns.forEach(btn => {
        btn.addEventListener('click', function(e) {
            e.preventDefault();
            const targetId = this.dataset.copy;
            const codeElement = document.getElementById(targetId);
            
            console.log('Copy button clicked:', targetId);
            
            if (codeElement) {
                const codeText = codeElement.textContent;
                
                // Copy to clipboard
                if (navigator.clipboard && navigator.clipboard.writeText) {
                    navigator.clipboard.writeText(codeText).then(() => {
                        console.log('Code copied successfully');
                        showCopyNotification('Code copied to clipboard!');
                        
                        // Update button text temporarily
                        const originalText = this.textContent;
                        this.textContent = 'Copied!';
                        setTimeout(() => {
                            this.textContent = originalText;
                        }, 1000);
                    }).catch(err => {
                        console.error('Failed to copy: ', err);
                        showCopyNotification('Failed to copy code', 'error');
                    });
                } else {
                    // Fallback for older browsers
                    try {
                        const textarea = document.createElement('textarea');
                        textarea.value = codeText;
                        document.body.appendChild(textarea);
                        textarea.select();
                        document.execCommand('copy');
                        document.body.removeChild(textarea);
                        showCopyNotification('Code copied to clipboard!');
                    } catch (err) {
                        console.error('Fallback copy failed:', err);
                        showCopyNotification('Copy not supported', 'error');
                    }
                }
            } else {
                console.error('Code element not found:', targetId);
            }
        });
    });
}

// Show copy notification - Fixed
function showCopyNotification(message, type = 'success') {
    // Remove existing notification
    const existingNotification = document.querySelector('.copy-notification');
    if (existingNotification) {
        existingNotification.remove();
    }
    
    // Create new notification
    const notification = document.createElement('div');
    notification.className = `copy-notification ${type}`;
    notification.textContent = message;
    document.body.appendChild(notification);
    
    // Show notification
    setTimeout(() => {
        notification.classList.add('show');
    }, 10);
    
    // Hide notification after 3 seconds
    setTimeout(() => {
        notification.classList.remove('show');
        setTimeout(() => {
            if (notification.parentNode) {
                notification.remove();
            }
        }, 300);
    }, 3000);
}

// Roadmap Progress Tracking - Fixed
function setupRoadmapProgress() {
    console.log('Setting up roadmap progress...');
    const markCompleteButtons = document.querySelectorAll('.mark-complete');
    
    console.log('Found mark complete buttons:', markCompleteButtons.length);
    
    markCompleteButtons.forEach(btn => {
        btn.addEventListener('click', function(e) {
            e.preventDefault();
            const topicId = this.dataset.topic;
            const topicCard = this.closest('.topic-card');
            const progressFill = topicCard.querySelector('.progress-fill');
            
            console.log('Mark complete clicked:', topicId);
            
            if (topicCard.classList.contains('completed')) {
                // Mark as incomplete
                topicCard.classList.remove('completed');
                progressFill.style.width = '0%';
                this.textContent = 'Mark Complete';
                delete roadmapProgress[topicId];
                console.log('Marked as incomplete:', topicId);
            } else {
                // Mark as complete
                topicCard.classList.add('completed');
                progressFill.style.width = '100%';
                this.textContent = 'Completed ✓';
                roadmapProgress[topicId] = true;
                console.log('Marked as complete:', topicId);
            }
            
            saveProgress();
        });
    });
}

// Exercise System - Fixed
function setupExercises() {
    console.log('Setting up exercises...');
    const optionBtns = document.querySelectorAll('.option-btn');
    
    console.log('Found option buttons:', optionBtns.length);
    
    optionBtns.forEach(btn => {
        btn.addEventListener('click', function(e) {
            e.preventDefault();
            const exerciseIndex = parseInt(this.dataset.exercise);
            const selectedOption = parseInt(this.dataset.option);
            const exercise = exercises[exerciseIndex];
            
            console.log('Exercise option clicked:', exerciseIndex, selectedOption);
            
            // Disable all options for this exercise
            const exerciseOptions = document.querySelectorAll(`[data-exercise="${exerciseIndex}"]`);
            exerciseOptions.forEach(option => {
                option.disabled = true;
                option.style.pointerEvents = 'none';
            });
            
            // Show feedback
            const feedbackElement = document.getElementById(`feedback-${exerciseIndex}`);
            
            if (selectedOption === exercise.correct) {
                // Correct answer
                this.classList.add('correct');
                feedbackElement.textContent = `Correct! ${exercise.explanation}`;
                feedbackElement.className = 'exercise-feedback correct';
                
                // Update score if not already completed
                if (!completedExercises.has(exerciseIndex)) {
                    currentScore++;
                    completedExercises.add(exerciseIndex);
                    updateScoreDisplay();
                    console.log('Exercise completed correctly, score:', currentScore);
                }
            } else {
                // Incorrect answer
                this.classList.add('incorrect');
                feedbackElement.textContent = `Incorrect. ${exercise.explanation}`;
                feedbackElement.className = 'exercise-feedback incorrect';
                
                // Highlight correct answer
                exerciseOptions[exercise.correct].classList.add('correct');
                console.log('Exercise answered incorrectly');
            }
            
            saveProgress();
        });
    });
}

// Update score display
function updateScoreDisplay() {
    const scoreElement = document.getElementById('score');
    if (scoreElement) {
        scoreElement.textContent = currentScore;
    }
}

// Reset exercises
function resetExercises() {
    console.log('Resetting exercises...');
    currentScore = 0;
    completedExercises.clear();
    
    // Reset all option buttons
    const optionBtns = document.querySelectorAll('.option-btn');
    optionBtns.forEach(btn => {
        btn.classList.remove('correct', 'incorrect');
        btn.disabled = false;
        btn.style.pointerEvents = 'auto';
    });
    
    // Clear all feedback
    const feedbackElements = document.querySelectorAll('.exercise-feedback');
    feedbackElements.forEach(feedback => {
        feedback.textContent = '';
        feedback.className = 'exercise-feedback';
    });
    
    updateScoreDisplay();
    saveProgress();
}

// Save progress to localStorage
function saveProgress() {
    const progressData = {
        currentScore: currentScore,
        completedExercises: Array.from(completedExercises),
        roadmapProgress: roadmapProgress
    };
    
    try {
        localStorage.setItem('goLearningProgress', JSON.stringify(progressData));
        console.log('Progress saved');
    } catch (e) {
        console.warn('Could not save progress:', e);
    }
}

// Load progress from localStorage
function loadProgress() {
    try {
        const saved = localStorage.getItem('goLearningProgress');
        if (saved) {
            const progressData = JSON.parse(saved);
            
            currentScore = progressData.currentScore || 0;
            completedExercises = new Set(progressData.completedExercises || []);
            roadmapProgress = progressData.roadmapProgress || {};
            
            updateScoreDisplay();
            restoreRoadmapProgress();
            restoreExerciseProgress();
            console.log('Progress loaded');
        }
        
        // Restore current tab
        const currentTab = localStorage.getItem('currentTab');
        if (currentTab) {
            const tabBtn = document.querySelector(`[data-tab="${currentTab}"]`);
            const tabContent = document.getElementById(currentTab);
            
            if (tabBtn && tabContent) {
                document.querySelectorAll('.tab-btn').forEach(btn => btn.classList.remove('active'));
                document.querySelectorAll('.tab-content').forEach(content => content.classList.remove('active'));
                
                tabBtn.classList.add('active');
                tabContent.classList.add('active');
                console.log('Restored tab:', currentTab);
            }
        }
    } catch (e) {
        console.warn('Could not load progress:', e);
    }
}

// Restore roadmap progress
function restoreRoadmapProgress() {
    Object.keys(roadmapProgress).forEach(topicId => {
        const topicCard = document.querySelector(`[data-topic="${topicId}"]`);
        if (topicCard) {
            const progressFill = topicCard.querySelector('.progress-fill');
            const markCompleteBtn = topicCard.querySelector('.mark-complete');
            
            topicCard.classList.add('completed');
            progressFill.style.width = '100%';
            markCompleteBtn.textContent = 'Completed ✓';
        }
    });
}

// Restore exercise progress
function restoreExerciseProgress() {
    completedExercises.forEach(exerciseIndex => {
        const exercise = exercises[exerciseIndex];
        const exerciseOptions = document.querySelectorAll(`[data-exercise="${exerciseIndex}"]`);
        const feedbackElement = document.getElementById(`feedback-${exerciseIndex}`);
        
        // Disable all options
        exerciseOptions.forEach(option => {
            option.disabled = true;
            option.style.pointerEvents = 'none';
        });
        
        // Show correct answer
        exerciseOptions[exercise.correct].classList.add('correct');
        
        // Show feedback
        feedbackElement.textContent = `Correct! ${exercise.explanation}`;
        feedbackElement.className = 'exercise-feedback correct';
    });
}

// Keyboard shortcuts
document.addEventListener('keydown', function(e) {
    // Ctrl/Cmd + number keys for tab navigation
    if ((e.ctrlKey || e.metaKey) && e.key >= '1' && e.key <= '5') {
        e.preventDefault();
        const tabIndex = parseInt(e.key) - 1;
        const tabBtns = document.querySelectorAll('.tab-btn');
        
        if (tabBtns[tabIndex]) {
            tabBtns[tabIndex].click();
        }
    }
    
    // Escape key to close expanded examples
    if (e.key === 'Escape') {
        const expandedHeaders = document.querySelectorAll('.example-header.expanded');
        expandedHeaders.forEach(header => {
            header.click();
        });
    }
});

// Export functions for global access
window.resetExercises = resetExercises;
window.goLearningApp = {
    resetExercises,
    saveProgress,
    loadProgress,
    currentScore,
    completedExercises,
    roadmapProgress
};

// Error handling
window.addEventListener('error', function(e) {
    console.error('Application error:', e.error);
    showCopyNotification('An error occurred. Please refresh the page.', 'error');
});

// Cleanup on page unload
window.addEventListener('beforeunload', function() {
    saveProgress();
});

console.log('Script loaded successfully');